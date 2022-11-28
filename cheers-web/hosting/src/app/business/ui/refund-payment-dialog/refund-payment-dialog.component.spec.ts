import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RefundPaymentDialogComponent } from './refund-payment-dialog.component';

describe('RefundPaymentDialogComponent', () => {
  let component: RefundPaymentDialogComponent;
  let fixture: ComponentFixture<RefundPaymentDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RefundPaymentDialogComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RefundPaymentDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
