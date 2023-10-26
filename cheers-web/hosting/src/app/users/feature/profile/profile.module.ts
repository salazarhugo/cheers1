import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {ProfileRoutingModule} from './profile-routing.module';
import {ProfileComponent} from "./profile.component";
import {PartyItemModule} from "../../../parties/ui/party-item/party-item.module";
import {ProfileHeaderModule} from "../../ui/profile-header/profile-header.module";
import {MatIconModule} from "@angular/material/icon";
import {FlexModule} from "@angular/flex-layout";
import {PostItemModule} from "../../../posts/ui/post-item/post-item.module";
import {MatTabsModule} from "@angular/material/tabs";
import {MatButtonModule} from "@angular/material/button";


@NgModule({
    declarations: [ProfileComponent],
    imports: [
        CommonModule,
        ProfileRoutingModule,
        MatTabsModule,
        PostItemModule,
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
