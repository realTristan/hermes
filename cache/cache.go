package cache

import "sync"

/*
The Cache struct is a data structure that represents a cache, which is a mechanism for temporarily
storing data in memory for faster access. The Cache struct contains a map that stores the cached
data, a pointer to a sync.RWMutex to provide thread-safe access to the cache, and a pointer to a
FullText struct for indexing and searching the cached data.

Fields:
@data: a map that stores the cached data in the form of a map of maps.
The outer map's keys represent the cache keys and the inner map's keys represent the field names of the cached data.

@mutex: a pointer to a sync.RWMutex to provide thread-safe access to the cache.
The mutex is used to ensure that multiple threads do not access the cache concurrently and cause race conditions.

@FT: a pointer to a FullText struct for indexing and searching the cached data.
The FullText struct is used to create a search index for the cached data to enable fast text-based searches.

Usage:

	cache := &Cache{
	  data: map[string]map[string]interface{}{},
	  mutex: &sync.RWMutex{},
	  FT: &FullText{},
	}
*/
type Cache struct {
	data  map[string]map[string]interface{} // The cache data.
	mutex *sync.RWMutex                     // Mutex for thread-safe access to the cache.
	FT    *FullText                         // FullText instance to store data in text format.
}

/*
*
The InitCache function is a factory function that creates and initializes a new Cache struct.
@Parameters: None
@Returns: A pointer to a newly created and initialized Cache struct.

Usage:

	cache := InitCache()

Example:

	cache := InitCache()
	// use the cache object for storing data
	cache.Set("key1", "field1", "value1")
	cache.Set("key2", "field2", "value2")
*/
func InitCache() *Cache {
	return &Cache{
		data:  map[string]map[string]interface{}{},
		mutex: &sync.RWMutex{},
		FT:    nil,
	}
}

/*
The Clean function is a method of the Cache struct that cleans the cache by deleting all the data from it.
It uses Mutex Locking to ensure that the cleaning process is thread-safe and can be accessed by only one thread at a time.

@Parameters: None
@Returns: None

Usage:

	cache := &Cache{}
	cache.Clean()

Example:

	cache := &Cache{}
	cache.Set("key1", "field1", "value1")
	cache.Set("key2", "field2", "value2")
	cache.Clean()

// The cache is now empty
*/
func (c *Cache) Clean() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.clean()
}

/*
The clean function is a private method of the Cache struct that is used to clear the cache by deleting all the data from it.
@Parameters: None
@Returns: None

Usage:

	This function is called internally by the Clean method of the Cache struct to clear the cache.
	It should not be called directly from outside the package.
*/
func (c *Cache) clean() {
	c.data = map[string]map[string]interface{}{}
}

/*
*

The Set function is a method of the Cache struct that sets a value in the cache for the specified key.
It uses Mutex Locking to ensure that the set operation is thread-safe and can be accessed by only one thread at a time.
@Parameters:

	key: a string that represents the key for the cached value.
	value: a map of string keys and interface{} values that represents the cached data.

@Returns: An error if there was an error setting the value in the cache, otherwise nil.

Usage:

	cache := &Cache{}
	err := cache.Set("key1", map[string]interface{}{"field1": "value1"})

Example:

	cache := &Cache{}
	err := cache.Set("key1", map[string]interface{}{"field1": "value1"})
	if err != nil {
		// handle error
	}
*/
func (c *Cache) Set(key string, value map[string]interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.set(key, value)
}

/*
The set function is a private method of the Cache struct that sets a value in the cache for the specified key.
@Parameters:

	key: a string that represents the key for the cached value.
	value: a map of string keys and interface{} values that represents the cached data.

@Returns: An error if there was an error setting the value in the cache, otherwise nil.

Usage:

	This function is called internally by the Set method of the Cache struct to set the value in the cache.
	It should not be called directly from outside the package.
*/
func (c *Cache) set(key string, value map[string]interface{}) error {
	// Update the value in the FT cache
	if c.FT != nil && c.FT.isInitialized {
		if err := c.setFT(key, value); err != nil {
			return err
		}
	}

	// Update the value in the cache
	c.data[key] = value

	// Return nil for no error
	return nil
}

/*
The Get function is a method of the Cache struct that retrieves a value from the cache for the specified key.
It uses RWMutex Locking to ensure that the get operation is thread-safe and can be accessed by multiple threads simultaneously.

@Parameters:

	key: a string that represents the key for the cached value.

@Returns: A map of string keys and interface{} values that represents the cached data for the specified key.

Usage:

	cache := &Cache{}
	value := cache.Get("key1")

Example:

	cache := &Cache{}
	err := cache.Set("key1", map[string]interface{}{"field1": "value1"})
	if err != nil {
		// handle error
	}
	value := cache.Get("key1")
*/
func (c *Cache) Get(key string) map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.get(key)
}

/*
This function retrieves the value associated with the given key from the cache.

@Parameters:

	key: string - The key used to retrieve the value from the cache.

@Returns:

	map[string]interface{}: The value associated with the given key. If the key is not present in the cache, this function returns nil.

Usage:

	This function is called to retrieve a value associated with a key from the cache.
	It does not perform any locking operations and is not thread-safe. It should only be called within a locked context,
	or within a thread-safe function such as Get().
*/
func (c *Cache) get(key string) map[string]interface{} {
	return c.data[key]
}

/*
Delete is a method of the Cache struct that deletes a key-value pair from the cache with Mutex Locking.

Usage:

	cache.Delete("mykey")

Parameters:

	key string: The key of the key-value pair to be deleted from the cache.

Returns:

	None.

The method first acquires a lock on the cache using the mutex.Lock() method, to ensure exclusive access to the cache.
The defer statement at the beginning of the method ensures that the lock is released after the method completes.
The delete(key) method is then called to remove the key-value pair from the cache.
Finally, the lock is released using the mutex.Unlock() method.
*/
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.delete(key)
}

/*
delete is a method of the Cache struct that removes a key-value pair from the cache.

Usage:

	c.Delete("mykey")

Parameters:

	key string: The key of the key-value pair to be removed from the cache.

Returns:

	None.

The method first checks if the cache's FT (Frequently Used Tags) feature is enabled and initialized,
and if so, deletes the key from the FT cache using the FT.Delete method. Next, it deletes the key-value
pair from the main cache using the built-in delete function.
*/
func (c *Cache) delete(key string) {
	// Delete the key from the FT cache
	if c.FT != nil && c.FT.isInitialized {
		c.deleteFT(key)
	}

	// Delete the key from the cache
	delete(c.data, key)
}

/*
Exists is a method of the Cache struct that checks if a key-value pair exists in the cache.

Usage:

	if myCache.Exists("mykey") {
		fmt.Println("Key exists in cache!")
	}

Parameters:

	key string: The key to check for in the cache.

Returns:

	bool: True if the key exists in the cache, false otherwise.

The method first acquires a read lock on the cache using the mutex.RLock() method, to ensure shared access to the cache.
The defer statement at the beginning of the method ensures that the lock is released after the method completes.
The exists(key) method is then called to check if the key exists in the cache.
Finally, the read lock is released using the mutex.RUnlock() method.

Example Usage:

	// Create a new cache instance
	myCache := InitCache()

	// Add some key-value pairs to the cache
	myCache.Set("key1", map[string]interface{}{"field1": "value1"})
	myCache.Set("key1", map[string]interface{}{"field1": "value2"})
	myCache.Set("key1", map[string]interface{}{"field1": "value3"})

	// Check if "key2" exists in the cache
	if myCache.Exists("key2") {
		fmt.Println("Key exists in cache!")
	} else {
		fmt.Println("Key does not exist in cache!")
	}

In this example, we first create a new Cache instance using the NewCache() function (not shown),
then add three key-value pairs to the cache using the Set method. Next, we check if one of the keys
exists in the cache using the Exists method with the key as its argument. Finally, we print a message
indicating whether the key exists in the cache or not.
*/
func (c *Cache) Exists(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.exists(key)
}

/*
exists is a method of the Cache struct that checks if a given key exists in the cache.

Usage:

	exists := myCache.Exists("mykey")

Parameters:

	key string: The key to check for in the cache.

Returns:

	bool: True if the key exists in the cache, false otherwise.

The method checks if the key exists in the map c.data, which stores the key-value pairs in the cache.
The map is accessed using the key in a comma-okay idiom to check if the key exists in the map.
If the key exists, the second value returned by the comma-okay idiom will be true, indicating that the key exists in the map.
Otherwise, the second value will be false, indicating that the key does not exist in the map.
*/
func (c *Cache) exists(key string) bool {
	_, ok := c.data[key]
	return ok
}

/*
Keys is a method of the Cache struct that returns an array of all the keys in the cache.

Usage:

	keys := myCache.Keys()

Returns:

	[]string: An array of strings representing all the keys in the cache.

The method returns a new array of strings containing all the keys in the cache.
The method does not modify the cache, so it is safe to call concurrently from multiple goroutines.
The keys are returned in an arbitrary order.

The method first acquires a read lock on the cache's mutex to ensure that it is safe to read from the cache.
Then, it calls the `keys` method to retrieve the keys and returns the result.
Finally, it releases the read lock on the mutex to allow other goroutines to access the cache.

Note that the keys array returned by this method is a copy of the keys in the cache at the time of the method call.
Therefore, changes to the cache made after the method call will not be reflected in the returned array.

Example Usage:

	// Create a new cache instance
	myCache := InitCache()

	// Add some key-value pairs to the cache
	myCache.Set("key1", map[string]interface{}{"field1": "value1"})
	myCache.Set("key1", map[string]interface{}{"field1": "value2"})
	myCache.Set("key1", map[string]interface{}{"field1": "value3"})

	// Get an array of all keys in the cache
	keys := myCache.Keys()

	// Print the keys in the cache
	fmt.Println("Keys in cache:", keys)

In this example, we first create a new Cache instance using the NewCache() function (not shown),
then add three key-value pairs to the cache using the Set method. Next, we call the Keys method
to retrieve an array of all the keys in the cache, and store the result in a variable called keys.
Finally, we print the keys in the cache using fmt.Println. Note that the order of the keys in the
printed output may differ from the order in which they were added to the cache, because the order
of keys returned by the Keys method is arbitrary.
*/
func (c *Cache) Keys() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.keys()
}

/*
keys is a private method of the Cache struct that returns an array of all the keys in the cache.

Usage:

	keys := c.keys()

Returns:

	[]string: An array of strings representing all the keys in the cache.

The method returns a new array of strings containing all the keys in the cache.
The method does not modify the cache, so it is safe to call concurrently from multiple goroutines.
The keys are returned in an arbitrary order.

The method iterates over the keys in the cache's `data` map and appends each key to an array.
Finally, it returns the resulting array of keys. This method should only be called within a
locked context or within a thread-safe function, as it accesses the `data` map of the cache directly.

Note that the keys array returned by this method is a copy of the keys in the cache at the time of the method call.
Therefore, changes to the cache made after the method call will not be reflected in the returned array.
*/
func (c *Cache) keys() []string {
	var keys []string = []string{}
	for key := range c.data {
		keys = append(keys, key)
	}
	return keys
}

/*
Values is a thread-safe method of the Cache struct that returns an array of all the values in the cache.

Usage:

	values := c.Values()

Returns:

	[]map[string]interface{}: An array of maps representing all the values in the cache.

The method returns a new array of maps containing all the values in the cache.
The method does not modify the cache, so it is safe to call concurrently from multiple goroutines.
The values are returned in an arbitrary order.

The method iterates over the values in the cache's `data` map and appends each value to an array.
Each value is represented as a map with a single key-value pair, where the key is the original key in the cache,
and the value is the cached value associated with that key.

Note that the values array returned by this method is a copy of the values in the cache at the time of the method call.
Therefore, changes to the cache made after the method call will not be reflected in the returned array.

Example Usage:

	// Create a new cache instance
	myCache := InitCache()

	// Add some key-value pairs to the cache
	myCache.Set("key1", map[string]interface{}{"field1": "value1"})
	myCache.Set("key1", map[string]interface{}{"field1": "value2"})
	myCache.Set("key1", map[string]interface{}{"field1": "value3"})

	// Get an array of all values in the cache
	values := myCache.Values()

	// Print the values in the cache
	fmt.Println("Values in cache:", values)

In this example, we first create a new Cache instance using the NewCache() function (not shown),
then add three key-value pairs to the cache using the Set method. Next, we call the Values method
to retrieve an array of all the values in the cache, and store the result in a variable called values.
Finally, we print the values in the cache using fmt.Println. Note that the order of the values in
the printed output may differ from the order in which they were added to the cache, because the
order of values returned by the Values method is arbitrary.
*/
func (c *Cache) Values() []map[string]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.values()
}

/*
Returns all the values stored in the cache.

@Parameters:
- c: a pointer to the Cache instance

@Returns:
- values: a slice of map[string]interface{}, where each map represents the value stored in the cache.

Example Usage:

	cache := NewCache()
	cache.Set("key1", map[string]interface{}{"name": "John", "age": 30})
	cache.Set("key2", map[string]interface{}{"name": "Jane", "age": 25})
	cache.Set("key3", map[string]interface{}{"name": "Bob", "age": 40})
	values := cache.values()
	fmt.Println(values)
*/
func (c *Cache) values() []map[string]interface{} {
	var values []map[string]interface{} = []map[string]interface{}{}
	for _, value := range c.data {
		values = append(values, value)
	}
	return values
}

/*
Returns the number of items stored in the cache using mutex locking to ensure thread-safety.

@Parameters:
  - c: a pointer to the Cache instance

@Returns:
  - length: an integer that represents the number of items stored in the cache.

Example Usage:

	cache := NewCache()
	cache.Set("key1", map[string]interface{}{"name": "John", "age": 30})
	cache.Set("key2", map[string]interface{}{"name": "Jane", "age": 25})
	cache.Set("key3", map[string]interface{}{"name": "Bob", "age": 40})
	length := cache.Length()
	fmt.Println(length)
*/
func (c *Cache) Length() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.length()
}

/*
Returns the number of items stored in the cache.

Parameters:
  - c: a pointer to the Cache instance

Returns:
  - length: an integer that represents the number of items stored in the cache.
*/
func (c *Cache) length() int {
	return len(c.data)
}
