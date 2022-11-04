import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {ProfileRoutingModule} from './profile-routing.module';
import {ProfileComponent} from "./profile.component";
import {MatTabsModule} from "@angular/material/tabs";
import {PostsModule} from "../../../posts/posts.module";
import {PartyItemModule} from "../../../parties/ui/party-item/party-item.module";
import {ProfileHeaderModule} from "../../ui/profile-header/profile-header.module";
import {MatIconModule} from "@angular/material/icon";
import {MatButtonModule} from "@angular/material/button";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
    declarations: [ProfileComponent],
    imports: [
        CommonModule,
        ProfileRoutingModule,
        MatTabsModule,
        PostsModule,
        PartyItemModule,
        ProfileHeaderModule,
        MatIconModule,
        MatButtonModule,
        FlexModule
    ],
    exports: [ProfileComponent],
})
export class ProfileModule {
}
