import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PartyItemComponent} from "./party-item.component";
import {MaterialModule} from "../../../material/material.module";
import {RouterModule} from "@angular/router";

@NgModule({
    declarations: [PartyItemComponent],
    imports: [
        CommonModule,
        MaterialModule,
        RouterModule,
    ],
    exports: [PartyItemComponent],
})
export class PartyItemModule {
}
