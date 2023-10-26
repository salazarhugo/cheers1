import {Component, Inject, OnInit} from '@angular/core';
import {MAT_LEGACY_DIALOG_DATA as MAT_DIALOG_DATA, MatLegacyDialogRef as MatDialogRef} from "@angular/material/legacy-dialog";

@Component({
    selector: 'app-post-delete-dialog',
    templateUrl: './post-delete-dialog.component.html',
    styleUrls: ['./post-delete-dialog.component.sass']
})
export class PostDeleteDialogComponent {

    constructor(
        public dialogRef: MatDialogRef<PostDeleteDialogComponent>,
        @Inject(MAT_DIALOG_DATA) public data: any,
    ) {
    }

    onNoClick(): void {
        this.dialogRef.close();
    }
}
