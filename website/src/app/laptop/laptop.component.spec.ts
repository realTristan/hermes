import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LaptopComponent } from './laptop.component';

describe('LaptopComponent', () => {
  let component: LaptopComponent;
  let fixture: ComponentFixture<LaptopComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [LaptopComponent]
    });
    fixture = TestBed.createComponent(LaptopComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
