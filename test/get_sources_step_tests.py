import unittest

from watergrid.context import DataContext
from watergrid.pipelines.pipeline import Pipeline
from watergrid.steps import Step

from app.steps.get_sources_step import GetSourcesStep


class VerifyStep(Step):
    def __init__(self):
        super().__init__(self.__class__.__name__, requires=['source'])
        self._flag = False

    def run(self, context: DataContext):
        if context.get('source') is not None:
            if context.get('source').get_url() == 'test' or context.get('source').get_url() == 'test2':
                self._flag = True

    def get_flag(self) -> bool:
        return self._flag


class GetSourcesStepTestCase(unittest.TestCase):
    def test_step_returns_sources(self):
        pipeline = Pipeline('test_pipeline')
        step1 = VerifyStep()
        pipeline.add_step(step1)
        pipeline.add_step(GetSourcesStep(['test', 'test2']))
        pipeline.run()
        self.assertTrue(step1.get_flag())


if __name__ == '__main__':
    unittest.main()
