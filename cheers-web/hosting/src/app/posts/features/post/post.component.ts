import {Component, Input, OnInit} from '@angular/core';
import {Post} from "../../../shared/data/models/post.model";
import {PostService} from "../../data/post.service";
import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";

@Component({
    selector: 'app-post[postResponse]',
    templateUrl: './post.component.html',
    styleUrls: ['./post.component.sass']
})
export class PostComponent implements OnInit {

    @Input() postResponse: PostResponse
    comment: string = ""

    constructor(
        private postService: PostService,
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

}
