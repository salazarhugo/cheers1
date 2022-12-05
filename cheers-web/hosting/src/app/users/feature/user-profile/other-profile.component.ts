import {Component, OnInit} from '@angular/core';
import {UserService} from "../../../shared/data/services/user.service";
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {firstValueFrom, Observable, of} from "rxjs";
import {User} from "../../../shared/data/models/user.model";
import {PostService} from "../../../posts/data/post.service";
import {Post} from "../../../shared/data/models/post.model";
import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {AuthService} from "../../../shared/data/services/auth.service";
import {ChatService} from "../../../chats/data/chat.service";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
    selector: 'app-user-profile',
    templateUrl: './other-profile.component.html',
    styleUrls: ['./other-profile.component.sass']
})
export class OtherProfileComponent implements OnInit {

    $user: Observable<User | null> = of(null)
    $posts: Observable<PostResponse[] | null> = of(null)
    username: string | null = null
    isAdmin: boolean = false

    constructor(
        private userService: UserService,
        private route: ActivatedRoute,
        private router: Router,
        private postService: PostService,
        private auth: AngularFireAuth,
        private authService: AuthService,
        private chatService: ChatService,
        private snackbar: MatSnackBar,
    ) {
    }

    async ngOnInit() {
        const token = await firstValueFrom(this.auth.idTokenResult)
        if (token == null)
            return
        this.isAdmin = token.claims["admin"] != null

        this.route.paramMap.subscribe((params: ParamMap) => {
            const username = params.get("username")
            this.username = username
            if (username) {
                this.$user = this.userService.getUser(username)
                this.$posts = this.postService.getUserPosts(username)
            }
        })
    }

    async onMessageClick() {
        const otherUser = await firstValueFrom(this.$user)
        try {
            const room = await firstValueFrom(this.chatService.createRoom([otherUser?.id!]))
            await this.router.navigate(['chats', room.id])
        } catch (e) {
            this.snackbar.open("Failed to create room")
        }
    }

    async unfollowUser(username: string) {
        await firstValueFrom(this.userService.unfollowUser(username))
    }

    async followUser(username: string) {
        await firstValueFrom(this.userService.followUser(username))
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }

    promoteToBusinessAccount(uid: string) {
        this.authService.promoteToBusiness(uid).subscribe()
    }

}
