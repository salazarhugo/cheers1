import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";

@Component({
  selector: 'app-refund-payment-dialog',
  templateUrl: './refund-payment-dialog.component.html',
  styleUrls: ['./refund-payment-dialog.component.sass']
})
export class RefundPaymentDialogComponent implements OnInit {

  constructor(
      public dialogRef: MatDialogRef<RefundPaymentDialogComponent>,
      @Inject(MAT_DIALOG_DATA) public data: any,
  ) { }

  ngOnInit(): void {
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}
