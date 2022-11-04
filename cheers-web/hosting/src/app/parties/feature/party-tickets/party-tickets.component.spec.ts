import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PartyTicketsComponent } from './party-tickets.component';

describe('PartyTicketsComponent', () => {
  let component: PartyTicketsComponent;
  let fixture: ComponentFixture<PartyTicketsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PartyTicketsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PartyTicketsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
