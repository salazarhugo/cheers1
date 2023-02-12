import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {PartyTransferComponent} from "./party-transfer.component";

const routes: Routes = [{path: '', component: PartyTransferComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PartyTransferRoutingModule { }
