import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {EditPartyComponent} from "./edit-party.component";

const routes: Routes = [{path: '', component: EditPartyComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PartyEditRoutingModule { }
