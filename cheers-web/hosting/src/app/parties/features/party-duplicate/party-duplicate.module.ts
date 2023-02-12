import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PartyDuplicateRoutingModule } from './party-duplicate-routing.module';
import { PartyDuplicateComponent } from './party-duplicate.component';


@NgModule({
  declarations: [
    PartyDuplicateComponent
  ],
  imports: [
    CommonModule,
    PartyDuplicateRoutingModule
  ]
})
export class PartyDuplicateModule { }
