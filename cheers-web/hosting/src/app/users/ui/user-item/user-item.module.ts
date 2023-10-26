import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {UserItemRoutingModule} from './user-item-routing.module';
import {UserItemComponent} from "./user-item.component";
import {FlexModule} from "@angular/flex-layout";
import {MatLegacyTooltipModule as MatTooltipModule} from "@angular/material/legacy-tooltip";


@NgModule({
    declarations: [UserItemComponent],
    imports: [
        CommonModule,
        UserItemRoutingModule,
        FlexModule,
        MatTooltipModule
    ],
    exports: [UserItemComponent],
})
export class UserItemModule {
}
