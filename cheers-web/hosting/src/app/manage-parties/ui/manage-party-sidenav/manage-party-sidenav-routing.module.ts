import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {ManagePartySidenavComponent} from "./manage-party-sidenav.component";

const routes: Routes = [{path: '', component: ManagePartySidenavComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ManagePartySidenavRoutingModule { }
