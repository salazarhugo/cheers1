import {Component, Inject, OnInit} from '@angular/core';
import {MAT_LEGACY_DIALOG_DATA as MAT_DIALOG_DATA, MatLegacyDialogRef as MatDialogRef} from "@angular/material/legacy-dialog";
import {UserService} from "../../../shared/data/services/user.service";

@Component({
  selector: 'app-party-duplicate',
  templateUrl: './party-duplicate.component.html',
  styleUrls: ['./party-duplicate.component.sass']
})
export class PartyDuplicateComponent implements OnInit {

  constructor(
      public dialogRef: MatDialogRef<PartyDuplicateComponent>,
      @Inject(MAT_DIALOG_DATA) public data: any,
  ) { }

  ngOnInit(): void {
  }

  onNoClick() {
    this.dialogRef.close(false);
  }
}
