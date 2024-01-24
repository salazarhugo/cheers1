import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PartyDuplicateRoutingModule } from './party-duplicate-routing.module';
import { PartyDuplicateComponent } from './party-duplicate.component';
import {MatDialogModule} from "@angular/material/dialog";
import {FlexModule} from "@angular/flex-layout";
import {MatIconModule} from "@angular/material/icon";
import {MatButtonModule} from "@angular/material/button";


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
