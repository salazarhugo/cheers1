import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PartyTransferComponent } from './party-transfer.component';

describe('PartyTransferComponent', () => {
  let component: PartyTransferComponent;
  let fixture: ComponentFixture<PartyTransferComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PartyTransferComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PartyTransferComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
