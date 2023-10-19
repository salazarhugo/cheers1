import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PartyTransferRoutingModule } from './party-transfer-routing.module';
import { PartyTransferComponent } from './party-transfer.component';
import {MatDialogModule} from "@angular/material/dialog";
import {MatButtonModule} from "@angular/material/button";
import {MatIconModule} from "@angular/material/icon";
import {MatMenuModule} from "@angular/material/menu";
import {FlexModule} from "@angular/flex-layout";
import {MatFormFieldModule} from "@angular/material/form-field";
import {UserItemModule} from "../../../users/ui/user-item/user-item.module";
import {SharedModule} from "../../../shared/shared.module";


@NgModule({
  declarations: [
    PartyTransferComponent
  ],
    imports: [
        CommonModule,
        PartyTransferRoutingModule,
        MatDialogModule,
        MatButtonModule,
        MatIconModule,
        MatMenuModule,
        FlexModule,
        MatFormFieldModule,
        UserItemModule,
        SharedModule
    ]
})
export class PartyTransferModule { }
