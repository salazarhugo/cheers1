import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyInviteRoutingModule} from './party-invite-routing.module';
import {PartyInviteComponent} from './party-invite.component';
import {MatDialogModule} from "@angular/material/dialog";
import {MatButtonModule} from "@angular/material/button";
import {MatChipsModule} from '@angular/material/chips';
import {UserItemModule} from "../../../users/ui/user-item/user-item.module";
import {MatRadioModule} from "@angular/material/radio";
import {MatCheckboxModule} from "@angular/material/checkbox";
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";
import {FlexModule} from "@angular/flex-layout";

@NgModule({
    declarations: [
        PartyInviteComponent
    ],
    imports: [
        CommonModule,
        PartyInviteRoutingModule,
        MatDialogModule,
        MatButtonModule,
        MatChipsModule,
        UserItemModule,
        MatRadioModule,
        MatCheckboxModule,
        MatProgressSpinnerModule,
        FlexModule,
    ],
    exports: [
        PartyInviteComponent
    ],
})
export class PartyInviteModule {
}
