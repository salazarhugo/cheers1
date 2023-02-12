import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {UserProfileRoutingModule} from './user-profile-routing.module';
import {OtherProfileComponent} from "./other-profile.component";
import {MatTabsModule} from "@angular/material/tabs";
import {ProfileHeaderModule} from "../../ui/profile-header/profile-header.module";
import {PostItemModule} from "../../../posts/ui/post-item/post-item.module";
import {MatButtonModule} from "@angular/material/button";
import {FlexModule} from "@angular/flex-layout";
import {MatIconModule} from "@angular/material/icon";
import {MatMenuModule} from "@angular/material/menu";
import {PartyItemModule} from "../../../parties/ui/party-item/party-item.module";


@NgModule({
    declarations: [OtherProfileComponent],
    imports: [
        CommonModule,
        UserProfileRoutingModule,
        MatTabsModule,
        PostItemModule,
        ProfileHeaderModule,
        MatButtonModule,
        FlexModule,
        MatIconModule,
        MatMenuModule,
        PartyItemModule,
    ],
    exports: [OtherProfileComponent],
})
export class UserProfileModule {
}
