import {User} from "../../../../gen/ts/cheers/type/user/user";
import {UserModel} from "../models/user.model";


export function toUserModel(u: User): UserModel {
    const res = new UserModel()
    res.id = u.id
    res.name = u.name
    res.username = u.username
    res.picture = u.picture
    res.verified = u.verified
    res.bio = u.bio
    res.email = u.email
    res.admin = u.isAdmin
    res.business = u.isBusinessAccount

    return res
}