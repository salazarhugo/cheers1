import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {UserResolver} from "./users/data/user.resolver";
import {CoreComponent} from "./core/feature/core.component";
import {SignInComponent} from "./shared/ui/sign-in/sign-in.component";
import {FinishSignUpComponent} from "./shared/ui/finish-sign-up/finish-sign-up.component";
import {TicketingComponent} from "./shared/ui/ticketing/ticketing.component";
import {PaymentComponent} from "./payments/features/payment/payment.component";
import {CompleteComponent} from "./shared/ui/complete/complete.component";
import {StoriesComponent} from "./stories/features/story-feed/stories.component";
import {ManagePartyComponent} from "./manage-parties/feature/manage-party/manage-party.component";
import {PartyResolver} from "./manage-parties/data/party.resolver";
import {ProfileResolver} from "./users/data/profile.resolver";
import {SignUpComponent} from "./sign-up/sign-up.component";

const routes: Routes = [
    {
        path: 'finish-sign-up',
        component: FinishSignUpComponent,
    },
    {
        path: 'payment/complete',
        component: CompleteComponent,
    },
    {
        path: 'payment/:clientSecret',
        component: PaymentComponent,
    },
    {
        path: 'story-feed/:id',
        component: StoriesComponent,
    },
    {
        path: 'ticketing/:id',
        component: TicketingComponent,
    },
    {
        path: 'events',
        redirectTo: 'parties',
    },
    {
        path: 'sign-in',
        component: SignInComponent,
    },
    {
        path: 'sign-up',
        component: SignUpComponent,
        loadChildren: () => import('./sign-up/sign-up.module').then(m => m.SignUpModule)
    },
    {
        path: 'manage/parties',
        component: ManagePartyComponent, resolve: {user: UserResolver, party: PartyResolver},
        loadChildren: () => import('./manage-parties/feature/manage-party-shell/manage-party-shell.module').then(m => m.ManagePartyShellModule)
    },
    {
        path: '',
        redirectTo: 'sign-in',
        pathMatch: 'full',
    },
    {
        path: '',
        component: CoreComponent,
        resolve: {data: UserResolver},
        loadChildren: () => import('./core/core.module').then(m => m.CoreModule)
    },
    // { path: '**', component: PageNotFoundComponent },  // Wildcard route for a 404 page
    // {
    //     path: '404',
    //     redirectTo: 'sign-in',
    // },
];

@NgModule({
    declarations: [],
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
