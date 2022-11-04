import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {UserProfileRoutingModule} from './user-profile-routing.module';
import {OtherProfileComponent} from "./other-profile.component";
import {MatTabsModule} from "@angular/material/tabs";
import {PostsModule} from "../../../posts/posts.module";
import {ProfileHeaderModule} from "../../ui/profile-header/profile-header.module";


@NgModule({
    declarations: [OtherProfileComponent],
    imports: [
        CommonModule,
        UserProfileRoutingModule,
        MatTabsModule,
        PostsModule,
        ProfileHeaderModule,
    ],
    exports: [OtherProfileComponent],
})
export class UserProfileModule {
}
