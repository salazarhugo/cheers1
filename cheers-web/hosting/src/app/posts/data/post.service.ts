import {Injectable} from '@angular/core';
import {ApiService} from "../../shared/data/services/api.service";
import {Post} from "../../shared/data/models/post.model";

@Injectable({
    providedIn: 'root'
})
export class PostService {

    constructor(
        private api: ApiService,
    ) {
    }

    createPost(post: Post) {
        return this.api.createPost(post)
    }

    likePost(postId: string) {
        this.api.likePost(postId).subscribe(res => console.log(res))
    }

    unlikePost(postId: string) {
        this.api.unlikePost(postId).subscribe(res => console.log(res))
    }

    getUserPosts(userId: string) {
        return this.api.getUserPosts(userId)
    }
}
