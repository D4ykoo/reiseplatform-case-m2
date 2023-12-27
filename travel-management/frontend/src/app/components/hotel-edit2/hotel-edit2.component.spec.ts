import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HotelEdit2Component } from './hotel-edit2.component';

describe('HotelEdit2Component', () => {
  let component: HotelEdit2Component;
  let fixture: ComponentFixture<HotelEdit2Component>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HotelEdit2Component]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(HotelEdit2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
