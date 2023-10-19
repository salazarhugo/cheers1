import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {TopbarComponent} from "./topbar.component";
import {MaterialModule} from "../../../material/material.module";
import {RouterModule} from "@angular/router";
import {UserChipModule} from "../../../users/ui/user-chip/user-chip.module";
import {FlexModule} from "@angular/flex-layout";
import {MatToolbarModule} from "@angular/material/toolbar";
import {MatMenuModule} from "@angular/material/menu";
import {MatIconModule} from "@angular/material/icon";
import {MatButtonModule} from "@angular/material/button";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatSlideToggleModule} from "@angular/material/slide-toggle";
import {MatDividerModule} from "@angular/material/divider";
import {MatTooltipModule} from "@angular/material/tooltip";
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
        MatToolbarModule,
        MatMenuModule,
        MatIconModule,
        MatButtonModule,
        MatFormFieldModule,
        MatSlideToggleModule,
        MatDividerModule,
        MatTooltipModule,
        UserItemModule,
    ],
    exports: [
        TopbarComponent,
    ]
})
export class TopbarModule {
}
