import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagePartyTicketsComponent } from './manage-party-tickets.component';

describe('ManagePartyTicketsComponent', () => {
  let component: ManagePartyTicketsComponent;
  let fixture: ComponentFixture<ManagePartyTicketsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagePartyTicketsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ManagePartyTicketsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
