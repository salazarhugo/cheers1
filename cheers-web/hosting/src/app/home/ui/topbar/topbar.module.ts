import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TopbarComponent} from "./topbar.component";
import {MaterialModule} from "../../../material/material.module";
import {RouterModule} from "@angular/router";
import {UserChipModule} from "../../../users/ui/user-chip/user-chip.module";
import {FlexModule} from "@angular/flex-layout";
import {UserItemModule} from "../../../users/ui/user-item/user-item.module";


@NgModule({
    declarations: [
        TopbarComponent,
    ],
    imports: [
        CommonModule,
        RouterModule,
        UserChipModule,
        FlexModule,
        UserItemModule,
        MaterialModule,
    ],
    exports: [
        TopbarComponent,
    ]
})
export class TopbarModule {
}
