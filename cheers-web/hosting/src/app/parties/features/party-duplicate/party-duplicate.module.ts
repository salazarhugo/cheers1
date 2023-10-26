import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PartyDuplicateRoutingModule } from './party-duplicate-routing.module';
import { PartyDuplicateComponent } from './party-duplicate.component';
import {MatLegacyDialogModule as MatDialogModule} from "@angular/material/legacy-dialog";
import {FlexModule} from "@angular/flex-layout";
import {MatIconModule} from "@angular/material/icon";
import {MatLegacyButtonModule as MatButtonModule} from "@angular/material/legacy-button";


@NgModule({
  declarations: [
    PartyDuplicateComponent
  ],
  imports: [
    CommonModule,
    PartyDuplicateRoutingModule,
    MatDialogModule,
    FlexModule,
    MatIconModule,
    MatButtonModule
  ]
})
export class PartyDuplicateModule { }
