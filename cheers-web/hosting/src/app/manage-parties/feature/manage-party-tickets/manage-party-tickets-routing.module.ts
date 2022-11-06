import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {ManagePartyTicketsComponent} from "./manage-party-tickets.component";

const routes: Routes = [{path: '', component: ManagePartyTicketsComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ManagePartyTicketsRoutingModule { }
