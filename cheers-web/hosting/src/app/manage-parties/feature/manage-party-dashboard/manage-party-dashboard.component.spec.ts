import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagePartyDashboardComponent } from './manage-party-dashboard.component';

describe('ManagePartyDashboardComponent', () => {
  let component: ManagePartyDashboardComponent;
  let fixture: ComponentFixture<ManagePartyDashboardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagePartyDashboardComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ManagePartyDashboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
