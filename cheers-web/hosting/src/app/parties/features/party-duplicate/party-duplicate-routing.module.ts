import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {PartyDuplicateComponent} from "./party-duplicate.component";

const routes: Routes = [{path: '', component: PartyDuplicateComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PartyDuplicateRoutingModule { }
