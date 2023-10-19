import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BusinessPartiesComponent } from './business-parties.component';

describe('BusinessPartiesComponent', () => {
  let component: BusinessPartiesComponent;
  let fixture: ComponentFixture<BusinessPartiesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BusinessPartiesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BusinessPartiesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
