import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {PartyHostingComponent} from "./party-hosting.component";

const routes: Routes = [{path: '', component: PartyHostingComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PartyHostingRoutingModule { }
