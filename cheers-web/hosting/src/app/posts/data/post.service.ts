import {Injectable} from '@angular/core';
import {ApiService} from "../../shared/data/services/api.service";
import {Post} from "../../shared/data/models/post.model";
import {Observable} from "rxjs";
import {PostResponse} from "../../../gen/ts/cheers/post/v1/post_service";
import {Empty} from "../../../gen/ts/google/protobuf/empty";

@Injectable({
    providedIn: 'root'
})
export class PostService {

    constructor(
        private api: ApiService,
    ) {
    }

    createPost(post: Post): Observable<PostResponse> {
        return this.api.createPost(post)
    }

    deletePost(postId: string): Observable<Empty> {
        return this.api.deletePost(postId)
    }

    getPostFeed(): Observable<PostResponse[]> {
        return this.api.getPostFeed()
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
