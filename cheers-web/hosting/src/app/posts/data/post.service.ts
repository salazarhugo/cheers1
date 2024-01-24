import {Injectable} from '@angular/core';
import {ApiService} from "../../shared/data/services/api.service";
import {PostModel} from "../../shared/data/models/post.model";
import {map, Observable} from "rxjs";
import {PostResponse} from "../../../gen/ts/cheers/post/v1/post_service";
import {Empty} from "../../../gen/ts/google/protobuf/empty";
import {toPostModel} from "../../shared/data/mappers/post.mapper";
import {Post} from "../../../gen/ts/cheers/type/post/post";

@Injectable({
    providedIn: 'root'
})
export class PostService {

    constructor(
        private api: ApiService,
    ) {
    }

    createPost(post: PostModel): Observable<PostModel> {
        return this.api.createPost(post)
            .pipe(map(post => toPostModel(post)))
    }

    deletePost(postId: string): Observable<Empty> {
        return this.api.deletePost(postId)
    }

    listMapPost(): Observable<PostResponse[]> {
        return this.api.listMapPost()
    }

    getPostFeed(): Observable<PostModel[]> {
        return this.api.getPostFeed()
            .pipe(map(postList => postList.map(post => toPostModel(post))))
    }

    likePost(postId: string) {
        this.api.likePost(postId).subscribe(res => console.log(res))
    }

    unlikePost(postId: string) {
        this.api.unlikePost(postId).subscribe(res => console.log(res))
    }

    getUserPosts(username: string): Observable<PostModel[]> {
        return this.api.getUserPosts(username)
            .pipe(map(postList => postList.map(post => toPostModel(post))))
    }
}
