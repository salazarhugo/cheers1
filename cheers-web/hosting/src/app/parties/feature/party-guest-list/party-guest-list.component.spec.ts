import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PartyGuestListComponent } from './party-guest-list.component';

describe('PartyGuestListComponent', () => {
  let component: PartyGuestListComponent;
  let fixture: ComponentFixture<PartyGuestListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PartyGuestListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PartyGuestListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
