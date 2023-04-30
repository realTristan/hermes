
/*
To do:





*/

// Get data from the cache function
export const Get = (key: string): any => {
    // GET
    return fetch(`${host}/get`)
    .then(res => res.json())
    .then(res => res);
}

// Set data in the cache function
export const Set = (key: string, value: map[string]interface{}): any => {
    // POST
    let value = base64encode(jsonEncode(value))
    return fetch(`${host}/set?key=${key}&value=${value}`)
    .then(res => res.json())
    .then(res => res);
}

// Delete key from the cache
export const Delete = (key: string): any => {
    // DELETE
    return fetch(`${host}/delete?key=${key}`)
    .then(res => res.json())
    .then(res => res);
}