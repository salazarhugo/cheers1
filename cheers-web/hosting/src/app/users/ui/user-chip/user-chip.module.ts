import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {UserChipRoutingModule} from './user-chip-routing.module';
import {UserChipComponent} from "./user-chip.component";
import {FlexModule} from "@angular/flex-layout";
import {MatRippleModule} from "@angular/material/core";


@NgModule({
    declarations: [UserChipComponent],
    imports: [
        CommonModule,
        UserChipRoutingModule,
        FlexModule,
        MatRippleModule
    ],
    exports: [UserChipComponent],
})
export class UserChipModule {
}
