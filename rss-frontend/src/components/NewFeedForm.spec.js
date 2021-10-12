import { mount } from '@cypress/vue'
import NewFeedForm from "@/components/NewFeedForm";

it ('renders new feed form', () => {
    mount(NewFeedForm, {

    })
    cy.contains('Add')
})