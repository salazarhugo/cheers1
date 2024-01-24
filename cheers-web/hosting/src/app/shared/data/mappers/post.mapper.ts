import {PostResponse} from "../../../../gen/ts/cheers/post/v1/post_service";
import {PostModel} from "../models/post.model";
import {toDrinkModel} from "./drink.mapper";
import {toUserModel} from "./user.mapper";
import {DrinkModel} from "../models/drink.model";


export function toPostModel(p: PostResponse): PostModel {
    const res = new PostModel()
    const post = p.post!!
    res.id = post.id
    res.caption = post.caption
    // res.drink = toDrinkModel(p.post?.drink)
    res.photos = post.photos
    res.isCreator = p.isCreator
    res.likeCount = p.likeCount
    res.commentCount = p.commentCount
    res.user = toUserModel(p.user!!)
    res.createTime = post.createTime
    res.drink = new DrinkModel()
    res.hasLiked = p.hasLiked

    return res
}