import {Component, OnInit} from '@angular/core';
import {UserService} from "../../../shared/data/services/user.service";
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {firstValueFrom, Observable, of} from "rxjs";
import {UserModel} from "../../../shared/data/models/user.model";
import {PostService} from "../../../posts/data/post.service";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {AuthService} from "../../../shared/data/services/auth.service";
import {ChatService} from "../../../chats/data/chat.service";
import {Party} from "../../../shared/data/models/party.model";
import {PartyService} from "../../../parties/data/party.service";
import {MatSnackBar} from "@angular/material/snack-bar";
import {PostModel} from "../../../shared/data/models/post.model";

@Component({
    selector: 'app-user-profile',
    templateUrl: './other-profile.component.html',
    styleUrls: ['./other-profile.component.sass']
})
export class OtherProfileComponent implements OnInit {

    $user: Observable<UserModel | null> = of(null)
    $posts: Observable<PostModel[] | null> = of(null)
    $parties: Observable<Party[] | null> = of(null)
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
        private partyService: PartyService,
        private snackbar: MatSnackBar,
    ) {
    }

    async ngOnInit() {
        this.userService.user$.subscribe(user => this.isAdmin = user.moderator)

        this.route.paramMap.subscribe((params: ParamMap) => {
            const username = params.get("username")
            this.username = username
            if (username) {
                this.$user = this.userService.getUser(username)
                this.$posts = this.postService.getUserPosts(username)
                this.$parties = this.partyService.getMyParties(username)
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

    verifyUser(userId: string) {
        this.authService.verifyUser(userId).subscribe()
    }
}
