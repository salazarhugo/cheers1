import {Component, OnInit} from '@angular/core';
import {UserService} from "../../../shared/data/services/user.service";
import {ActivatedRoute, ParamMap} from "@angular/router";
import {Observable, of} from "rxjs";
import {User} from "../../../shared/data/models/user.model";
import {PostService} from "../../../posts/data/post.service";
import {Post} from "../../../shared/data/models/post.model";
import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";

@Component({
    selector: 'app-user-profile',
    templateUrl: './other-profile.component.html',
    styleUrls: ['./other-profile.component.sass']
})
export class OtherProfileComponent implements OnInit {

    $user: Observable<User | null> = of(null)
    $posts: Observable<PostResponse[] | null> = of(null)
    username: string | null = null

    constructor(
        private userService: UserService,
        private route: ActivatedRoute,
        private postService: PostService,
    ) {
    }


    ngOnInit(): void {
        this.route.paramMap.subscribe((params: ParamMap) => {
            const username = params.get("username")
            this.username = username
            if (username) {
                this.$user = this.userService.getUser(username)
                this.$posts = this.postService.getUserPosts(username)
            }
        })
    }

    unfollowUser(username: string) {
        this.userService.unfollowUser(username)
    }

    followUser(username: string) {
        this.userService.followUser(username)
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }
}
