import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagePartyGuestListComponent } from './manage-party-guest-list.component';

describe('ManagePartyGuestListComponent', () => {
  let component: ManagePartyGuestListComponent;
  let fixture: ComponentFixture<ManagePartyGuestListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagePartyGuestListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ManagePartyGuestListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
