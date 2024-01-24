import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {ProfileHeaderRoutingModule} from './profile-header-routing.module';
import {ProfileHeaderComponent} from "./profile-header.component";
import {FlexLayoutModule, FlexModule} from "@angular/flex-layout";
import {MatTooltipModule} from "@angular/material/tooltip";


@NgModule({
    declarations: [ProfileHeaderComponent],
    imports: [
        CommonModule,
        ProfileHeaderRoutingModule,
        FlexModule,
        FlexLayoutModule,
        MatTooltipModule
    ],
    exports: [ProfileHeaderComponent],
})
export class ProfileHeaderModule {
}
