import {Component, Inject, OnInit} from '@angular/core';
import {MAT_LEGACY_DIALOG_DATA as MAT_DIALOG_DATA, MatLegacyDialogRef as MatDialogRef} from "@angular/material/legacy-dialog";

@Component({
    selector: 'app-post-create-dialog',
    templateUrl: './post-create-dialog.component.html',
    styleUrls: ['./post-create-dialog.component.sass']
})
export class PostCreateDialogComponent implements OnInit {

    constructor(
        public dialogRef: MatDialogRef<PostCreateDialogComponent>,
        @Inject(MAT_DIALOG_DATA) public data: any,
    ) {
    }

    files: File[] = []

    onNoClick(): void {
        this.dialogRef.close();
    }

    ngOnInit(): void {
    }

    async onSelect(event: any) {
        this.files = event.addedFiles
    }

    onRemove(event: any) {
    }
}
