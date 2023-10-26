import {Component, Inject, OnInit} from '@angular/core';
import {MAT_LEGACY_DIALOG_DATA as MAT_DIALOG_DATA, MatLegacyDialogRef as MatDialogRef} from "@angular/material/legacy-dialog";

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
