import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TravelOfferEditComponent } from './travel-offer-edit.component';

describe('TravelOfferEditComponent', () => {
  let component: TravelOfferEditComponent;
  let fixture: ComponentFixture<TravelOfferEditComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TravelOfferEditComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(TravelOfferEditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
