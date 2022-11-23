import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BusinessPayoutsComponent } from './business-payouts.component';

describe('BusinessPayoutsComponent', () => {
  let component: BusinessPayoutsComponent;
  let fixture: ComponentFixture<BusinessPayoutsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BusinessPayoutsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BusinessPayoutsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
