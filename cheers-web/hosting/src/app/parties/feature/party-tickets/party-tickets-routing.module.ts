import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {PartyTicketsComponent} from "./party-tickets.component";

const routes: Routes = [{path: '', component: PartyTicketsComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PartyTicketsRoutingModule { }
