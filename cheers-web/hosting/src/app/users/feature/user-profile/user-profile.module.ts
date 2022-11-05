import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {UserProfileRoutingModule} from './user-profile-routing.module';
import {OtherProfileComponent} from "./other-profile.component";
import {MatTabsModule} from "@angular/material/tabs";
import {ProfileHeaderModule} from "../../ui/profile-header/profile-header.module";
import {PostItemModule} from "../../../posts/ui/post-item/post-item.module";


@NgModule({
    declarations: [OtherProfileComponent],
    imports: [
        CommonModule,
        UserProfileRoutingModule,
        MatTabsModule,
        PostItemModule,
        ProfileHeaderModule,
    ],
    exports: [OtherProfileComponent],
})
export class UserProfileModule {
}
