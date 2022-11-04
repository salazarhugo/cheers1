import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {OtherProfileComponent} from "../user-profile/other-profile.component";

const routes: Routes = [
    {
        path: 'profile',
        loadChildren: () => import('../profile/profile.module').then(m => m.ProfileModule),
    },
    {
        path: 'edit',
        loadChildren: () => import('../user-edit/edit.module').then(m => m.EditModule),
    },
    {
        path: ':username',
        loadChildren: () => import('../user-profile/user-profile.module').then(m => m.UserProfileModule),
    },
    {
        path: ':username/followers',
        loadChildren: () => import('../user-followers/user-followers.module').then(m => m.UserFollowersModule),
    },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class UserShellRoutingModule {
}
