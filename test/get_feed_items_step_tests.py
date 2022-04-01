import unittest

from watergrid.pipelines.pipeline import Pipeline
from watergrid.steps import Step

from app.steps.get_feed_items_step import GetFeedItemsStep
from app.steps.get_sources_step import GetSourcesStep


class VerifyStep(Step):
    def __init__(self):
        super().__init__(self.__class__.__name__, requires=['item'])
        self._flag = False

    def run(self, context):
        if context.get('item').title == "Item 1":
            self._flag = True

    def get_flag(self):
        return self._flag


class GetFeedItemsStepTestCase(unittest.TestCase):
    def test_gets_items(self):
        pipeline = Pipeline('test_pipeline')
        pipeline.add_step(GetSourcesStep(['test/rss_sample.xml']))
        pipeline.add_step(GetFeedItemsStep())
        verify_step = VerifyStep()
        pipeline.add_step(verify_step)
        pipeline.run()
        self.assertTrue(verify_step.get_flag())


if __name__ == '__main__':
    unittest.main()
