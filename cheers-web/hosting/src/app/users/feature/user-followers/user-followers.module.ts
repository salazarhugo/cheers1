import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {UserFollowersRoutingModule} from './user-followers-routing.module';
import {UserFollowersComponent} from './user-followers.component';
import {MatTabsModule} from "@angular/material/tabs";
import {UserItemModule} from "../../ui/user-item/user-item.module";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
    declarations: [
        UserFollowersComponent
    ],
    imports: [
        CommonModule,
        UserFollowersRoutingModule,
        MatTabsModule,
        UserItemModule,
        FlexModule
    ],
    exports: [
        UserFollowersComponent
    ],
})
export class UserFollowersModule {
}
