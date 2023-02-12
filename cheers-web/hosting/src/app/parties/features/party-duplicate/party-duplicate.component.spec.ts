import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PartyDuplicateComponent } from './party-duplicate.component';

describe('PartyDuplicateComponent', () => {
  let component: PartyDuplicateComponent;
  let fixture: ComponentFixture<PartyDuplicateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PartyDuplicateComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PartyDuplicateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
