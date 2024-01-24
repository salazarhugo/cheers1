import {Component, EventEmitter, Inject, Input, OnInit, Output} from '@angular/core';
import {PostService} from "../../data/post.service";
import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";
import {PostDeleteDialogComponent} from "../post-delete-dialog/post-delete-dialog.component";
import {MatDialog} from "@angular/material/dialog";
import {MatSnackBar} from "@angular/material/snack-bar";
import {PostModel} from "../../../shared/data/models/post.model";

@Component({
    selector: 'app-post-item[post]',
    templateUrl: './post-item.component.html',
    styleUrls: ['./post-item.component.sass']
})
export class PostItemComponent implements OnInit {

    @Input() post: PostModel
    @Output() onDelete = new EventEmitter<PostModel>()

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
        if (this.post.hasLiked) {
            this.postService.unlikePost(this.post.id!)
            this.post.hasLiked = false
            this.post.likeCount--
        } else {
            this.postService.likePost(this.post.id!)
            this.post.hasLiked = true
            this.post.likeCount++
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
            if (this.post == undefined) return

            if (result == true)
                this.postService.deletePost(this.post.id).subscribe(res => {
                    this.onDelete.emit(this.post)
                    this.openSnackBar("Your post was deleted", "Hide")
                })
        });
    }
}