import {UserModel} from "./user.model";
import {DrinkModel} from "./drink.model";


export class PostModel {
      user: UserModel;
      likeCount: number;
      commentCount: number;
      hasLiked: boolean;
      isCreator: boolean;

      id: string = ""
      caption: string = ""
      username: string = ""
      createTime: number = 0
      relativeTime: string = ""
      drink: DrinkModel;
      privacy: string = "FRIENDS"
      photos: string[] = []
      videoUrl: string = ""
      videoThumbnailUrl: string = ""
      drunkenness: number = 0
      latitude: number = 0.0
      longitude: number = 0.0
      locationName: string = ""
      allowJoin: boolean = true
      tagUsersId: string[] = []
      type: string = ""
      accountId: string = ""
}
