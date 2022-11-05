import {Component, EventEmitter, Inject, Input, OnInit, Output} from '@angular/core';
import {Post} from "../../../shared/data/models/post.model";
import {PostService} from "../../data/post.service";
import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";
import {MAT_DIALOG_DATA, MatDialog, MatDialogRef} from "@angular/material/dialog";
import {PostDeleteDialogComponent} from "../post-delete-dialog/post-delete-dialog.component";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
    selector: 'app-post-item[postResponse]',
    templateUrl: './post-item.component.html',
    styleUrls: ['./post-item.component.sass']
})
export class PostItemComponent implements OnInit {

    @Input() postResponse: PostResponse
    @Output() onDelete = new EventEmitter<PostResponse>()

    comment: string = ""

    constructor(
        public dialog: MatDialog,
        private postService: PostService,
        private _snackBar: MatSnackBar
    ) {
    }

    ngOnInit(): void {
    }

    onToggleLike() {
        if (this.postResponse.hasLiked) {
            this.postService.unlikePost(this.postResponse.post?.id!)
            this.postResponse.hasLiked = false
        } else {
            this.postService.likePost(this.postResponse.post?.id!)
            this.postResponse.hasLiked = true
        }
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }

    openSnackBar(message: string, action: string) {
        this._snackBar.open(message, action);
    }

    delete() {
        const dialogRef = this.dialog.open(PostDeleteDialogComponent, {
            panelClass: 'cheers-dialog'
        });

        dialogRef.afterClosed().subscribe(result => {
            if (this.postResponse.post == undefined) return

            if (result == true)
                this.postService.deletePost(this.postResponse.post.id).subscribe(res => {
                    this.onDelete.emit(this.postResponse)
                    this.openSnackBar("Your post was deleted", "Hide")
                })
        });
    }
}